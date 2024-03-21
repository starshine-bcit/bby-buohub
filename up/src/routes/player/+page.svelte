
<script lang="ts">
    /// <reference types="video.js" />

    export let data;
    
    import { onMount } from 'svelte';
	import videojs from 'video.js';
    let player: any

    
    onMount(() => {
      player = videojs('dashVideo');
    
      player.ready(() => {
        player.src({
          src: data.urlArray[5].videoUrl,
          type: 'application/dash+xml'
        });
    
        player.play();
      });
    
      return () => {
        if (player) {
          player.dispose();
        }
      };
    });
  </script>
  

  <div class="container mt-4">
    <div class="mt-4 p-5 bg-primary text-white rounded">
      <h1>Hello there!</h1>
      <p>BbyBuoHub is a dockerized video streaming platform</p>
      <!-- svelte-ignore a11y-media-has-caption -->
      <video id="dashVideo" width="800" height="600" class="video-js vjs-default-skin" controls poster={data.urlArray[5].url}></video>
    </div>
  </div>
  