<script lang="ts">
import { Login, login } from "./http";
import { hashPass } from "./hashing";

let ok = false;
let pass = "";
let result = "";

async function handleSubmit() {
	let loginRes = await login({ hash: hashPass(pass) })
	if (typeof loginRes === "string") {
		console.log(loginRes)
		result = "not accepted"
	}
	else if (loginRes) {
		result = "accepted";
	}
}

</script>

<h1>Client program</h1>

<form on:submit|preventDefault={handleSubmit}>
	<input type="text" bind:value={pass}>
	<button type="submit">
		Submit
	</button>
</form>

<p>Login result: {result} </p> 

<style>
	h1 {
		color: green;
	}
</style>
