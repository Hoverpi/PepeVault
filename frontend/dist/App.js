import 'svelte/internal/disclose-version';
import 'svelte/internal/flags/legacy';
import * as $ from 'svelte/internal/client';
import Hello from './Hello.js';

var root = $.from_html(`<main class="svelte-1n46o8q"><h1>Minimal Svelte app</h1> <input placeholder="your name" class="svelte-1n46o8q"/> <!></main>`);

export default function App($$anchor) {
	let name = $.mutable_source('Enrique');

	function onChange(e) {
		$.set(name, e.target.value);
	}

	var main = root();
	var input = $.sibling($.child(main), 2);

	$.remove_input_defaults(input);

	var node = $.sibling(input, 2);

	Hello(node, {
		get name() {
			return $.get(name);
		}
	});

	$.reset(main);
	$.template_effect(() => $.set_value(input, $.get(name)));
	$.event('input', input, onChange);
	$.append($$anchor, main);
}