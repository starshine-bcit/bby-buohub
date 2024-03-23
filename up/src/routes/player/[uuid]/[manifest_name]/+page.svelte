<script lang="ts">
	// @ts-nocheck
	import { onMount } from 'svelte';
	import videojs from 'video.js';
	export let data;

	let player: any;

	onMount(() => {
		player = videojs('dashVideo');

		player.ready(() => {
			player.src({
				src: `https://${data.CDN_HOST}/stream/${data.uuid}/${data.manifest_name}`,
				type: 'application/dash+xml'
			});

			player.play();
		});
	});
</script>

<svelte:head>
	<script
		src="https://cdnjs.cloudflare.com/ajax/libs/video.js/8.11.2/video.min.js"
		integrity="sha512-habJfo2HgYsF+7SBP3dr03wtYz3sclQZCws+I822LbtjqXhh2ByGyTHVHbItmL811+eEbcDIcsq86uhoa1MV4g=="
		crossorigin="anonymous"
		referrerpolicy="no-referrer"
	></script>
	<script
		src="https://cdnjs.cloudflare.com/ajax/libs/dashjs/4.7.3/dash.all.min.js"
		integrity="sha512-tTDH+TGRdveBPX5nFyp7PWf7AYArZ4qnOcrxzNpYckqE2oNjG/B4DltycGwFPs5Fngus6dYu7d158Mo/H+9f1A=="
		crossorigin="anonymous"
		referrerpolicy="no-referrer"
	></script>
	<script
		src="https://cdnjs.cloudflare.com/ajax/libs/videojs-contrib-dash/5.1.1/videojs-dash.min.js"
		integrity="sha512-jmpCwJ7o9/MxR36dZX+SQc4Ta2PDvMwM5MHmW0oDcy/UzkuppIj+F9FiN+UW/Q8adlOwb5Tx06zHsY/7yg4OYg=="
		crossorigin="anonymous"
		referrerpolicy="no-referrer"
	></script>

	<link href="https://vjs.zencdn.net/8.10.0/video-js.css" rel="stylesheet" />
</svelte:head>

<div class="container mt-4">
	<video
		id="dashVideo"
		width="800"
		height="600"
		class="video-js vjs-default-skin"
		controls
		aria-controls=""><track kind="captions" /></video
	>
</div>
