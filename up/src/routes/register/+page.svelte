<script lang="ts">
	import { enhance } from '$app/forms';
	export let form;
</script>

<h1>Register</h1>

<section>
	<form method="post" use:enhance>
		<div class="group">
			<label for="username">Username</label>
			<input type="username" name="username" id="username" required />
		</div>

		<div class="group">
			<label for="password">Password</label>
			<input type="password" name="password" id="password" required />
		</div>

		<div class="submit-container">
			<button type="submit">Login</button>
		</div>
		{#if form?.err}
			<div class="notice error">
				{form.err}
			</div>
		{/if}
	</form>

	<div class="actions">
		<a href="/login">Login</a>
	</div>
</section>
