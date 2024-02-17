<script lang="ts">
	import { enhance } from '$app/forms';
	export let form;
</script>

<section>
	<div class="container mt-4">
		<div class="card mx-auto" style="width: 25rem;">
			<div class="card-header bg-dark text-warning">Upload a video</div>
			<div class="card-body">
				<form method="post" use:enhance enctype="multipart/form-data">
					<div class="group mb-3">
						<label for="title" class="form-label">Title</label>
						<input type="text" name="title" id="title" class="form-control" required />
					</div>

					<div class="group mb-3">
						<label for="description" class="form-label">Description</label>
						<input type="text" name="description" id="description" class="form-control" required />
					</div>
					<div class="group mb-3">
						<label for="file" class="form-label">File</label>
						<input
							type="file"
							name="file"
							id="file"
							class="form-control"
							accept=".mkv, .mp4, .webm, .mpeg4, .avi"
							required
						/>
					</div>

					<div class="submit-container mb-3">
						<button type="submit" class="btn btn-primary">Submit</button>
					</div>

					{#if form?.err}
						<div class="notice error mt-3">
							{form.err}
						</div>
					{/if}
				</form>
			</div>
		</div>
	</div>
</section>
