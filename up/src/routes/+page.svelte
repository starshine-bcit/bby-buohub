<script>
	export let data;
</script>

<h1>Videos</h1>

<div class="image-grid">
	{#each data.urlArray as video}
		<div class="image-container">
			<h3>{video.title}</h3>
			<p>Description: {video.description}</p>
			<p>
				<a href="/player/{video.uuid}/{video.manifest_name}">
					<img src={video.url} alt={video.title} />
				</a>
			</p>
		</div>
	{/each}
</div>

<style>
	/* Flexbox grid styles */
	.image-grid {
		display: flex;
		flex-wrap: wrap;
		justify-content: space-between;
		margin-top: 20px;
	}

	.image-container {
		width: calc(33.33% - 10px);
		margin-bottom: 20px;
		text-align: center;
	}

	.image-container img {
		max-width: 100%;
		height: auto;
		border-radius: 5px;
	}
</style>
