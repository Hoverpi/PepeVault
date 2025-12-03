import 'svelte/internal/disclose-version';
import 'svelte/internal/flags/legacy';
import * as $ from 'svelte/internal/client';

var root = $.from_html(`<p class="svelte-b59zkr"> </p>`);

export default function Hello($$anchor, $$props) {
	let name = $.prop($$props, 'name', 8, 'world');
	var p = root();
	var text = $.child(p);

	$.reset(p);
	$.template_effect(() => $.set_text(text, `Hello ${name() ?? ''}!`));
	$.append($$anchor, p);
}