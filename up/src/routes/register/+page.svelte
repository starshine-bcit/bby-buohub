<script lang="ts">
	import { enhance } from '$app/forms';
	export let form;
</script>

<h1>Login</h1>

<section>
	<form method="post" use:enhance>
		<div class="group">
			<label for="email">Email</label>
			<input type="email" name="email" id="email" required />
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
		<a href="/signup">Sign Up</a>
	</div>
</section>
