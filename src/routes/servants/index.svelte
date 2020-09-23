<script context="module">
	export function preload({ params, query }) {
		return this.fetch(`/data/servants.json`).then(r => r.json()).then(servants => {
			return { servants };
		});
	}
</script>

<script>
	export let servants;
</script>

<style>
	table {
		margin: 0 0 1em 0;
		line-height: 1.5;
	}
</style>

<svelte:head>
	<title>Servants</title>
</svelte:head>

<h1>Servants</h1>

<table>
	{#each servants as servant}
	    <tr>
			<td>{servant.ID}</td>
			<td><img src={servant.Icon.Src}  alt={servant.Name.EN}/></td>
			<td>{servant.Name.JP}</td>
		</tr>
	{/each}
</table>