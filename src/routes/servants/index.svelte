<script context="module">
	export function preload({ params, query }) {
		return this.fetch(`/data/servants.json`).then(r => r.json()).then(servants => {
			return { servants };
		});
	}
</script>

<script>
	import ServantTable from "../../components/ServantTable.svelte";
	export let servants;
</script>

<svelte:head>
	<title>Servants</title>
</svelte:head>

<h1>Servants</h1>

<ServantTable servants={servants} />