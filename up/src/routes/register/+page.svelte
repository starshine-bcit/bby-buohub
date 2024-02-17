<script lang="ts">
	import { enhance } from '$app/forms';
	export let form;
</script>

<section>
	<div class="container mt-4">
		<div class="card mx-auto" style="width: 18rem;">
			<div class="card-header bg-dark text-warning">Register with bby-buohub</div>
			<div class="card-body">
				<form method="post" use:enhance>
					<div class="group mb-3">
						<label for="username" class="form-label">Username</label>
						<input type="username" name="username" id="username" class="form-control" required />
					</div>

					<div class="group mb-3">
						<label for="password" class="form-label">Password</label>
						<input type="password" name="password" id="password" class="form-control" required />
					</div>

					<div class="submit-container">
						<button type="submit" class="btn btn-primary">Register</button>
					</div>

					{#if form?.err}
						<div class="notice error">
							{form.err}
						</div>
					{/if}
				</form>
			</div>
		</div>
	</div>
</section>
