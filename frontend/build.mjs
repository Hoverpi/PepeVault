// build.mjs — compile every .svelte in src/ into dist/ as ESM (.js),
// convert imports from ".svelte" -> ".js", and concatenate CSS.
// Usage: node build.mjs
// Optional: node build.mjs --watch

import fs from 'fs/promises';
import path from 'path';
import { compile } from 'svelte/compiler';
import { argv } from 'process';

const SRC = 'src';
const DIST = 'dist';
const ENTRY = path.join(SRC, 'App.svelte'); // your entry component

// recursively read .svelte files
async function walk(dir) {
  const res = [];
  for (const entry of await fs.readdir(dir, { withFileTypes: true })) {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory()) {
      res.push(...(await walk(full)));
    } else if (entry.isFile() && entry.name.endsWith('.svelte')) {
      res.push(full);
    }
  }
  return res;
}

function svelteImportToJs(code) {
  // replace import specifiers that end with .svelte -> .js
  // naive but practical: replace occurrences of ".svelte" inside import/export/dynamic import specifiers
  return code.replace(/(\.svelte)(['"])/g, '.js$2');
}

async function compileAll() {
  await fs.rm(DIST, { recursive: true, force: true });
  await fs.mkdir(DIST, { recursive: true });

  const files = await walk(SRC);
  if (!files.length) throw new Error('No .svelte files found in src/');

  let combinedCss = '';

  for (const file of files) {
    const rel = path.relative(SRC, file);         // e.g. "App.svelte" or "sub/Comp.svelte"
    const outRel = rel.replace(/\.svelte$/, '.js'); // e.g. "App.js"
    const outPath = path.join(DIST, outRel);
    const outDir = path.dirname(outPath);
    await fs.mkdir(outDir, { recursive: true });

    const source = await fs.readFile(file, 'utf8');

    // compile to ESM (Svelte v4/v5: default output is ESM; 'format' option removed)
    const { js, css, warnings } = compile(source, {
      filename: file,
      generate: 'client',
      // no 'format' option (compiler produces ESM)
      // css: 'external' -> returns css.code in css
      css: 'external'
    });

    if (warnings && warnings.length) {
      for (const w of warnings) console.warn('Svelte warning:', w);
    }

    // Fix imports referencing other .svelte files so they point to .js files
    const transformedJs = svelteImportToJs(js.code);

    // write compiled JS module to dist (ESM)
    await fs.writeFile(outPath, transformedJs, 'utf8');

    // collect css
    if (css && css.code) {
      combinedCss += `/* from ${rel} */\n` + css.code + '\n';
    }
  }

  // write combined CSS to dist/app.css (create an empty file if no CSS)
  await fs.writeFile(path.join(DIST, 'app.css'), combinedCss || '', 'utf8');

  // create a small main.js bootstrap that imports the entry (App.js) and instantiates it
  const entryJsRel = path.relative(SRC, ENTRY).replace(/\.svelte$/, '.js'); // usually "App.js"
  const mainJs = `import App from './${entryJsRel}';

const app = new App({
  target: document.getElementById('app'),
  // props: {}
});

export default app;
`;
  await fs.writeFile(path.join(DIST, 'main.js'), mainJs, 'utf8');

  // write index.html referencing module main.js
  const index = `<!doctype html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width,initial-scale=1" />
    <title>Svelte ESM no-bundle</title>
    <link rel="stylesheet" href="/app.css" />
  </head>
  <body>
    <div id="app"></div>
    <script type="module" src="/main.js"></script>
  </body>
</html>`;
  await fs.writeFile(path.join(DIST, 'index.html'), index, 'utf8');

  console.log('Built ESM modules into', DIST);
}

if (argv.includes('--watch')) {
  console.log('Watch mode — recompiling on changes...');
  let timeout = null;
  const rebuild = async () => {
    clearTimeout(timeout);
    timeout = setTimeout(async () => {
      try {
        await compileAll();
      } catch (e) {
        console.error('Build failed:', e);
      }
    }, 120);
  };
  // initial build
  rebuild();
  // fs.watch the src tree
  const chokify = async () => {
    const watcher = require('fs').watch(SRC, { recursive: true }, () => rebuild());
    process.on('exit', () => watcher.close());
  };
  chokify();
} else {
  compileAll().catch(e => {
    console.error('Build failed:', e);
    process.exit(1);
  });
}
