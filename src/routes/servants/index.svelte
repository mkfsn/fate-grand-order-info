<script context="module">
	export function preload({ params, query }) {
		return this.fetch(`/data/servants.json`).then(r => r.json()).then(servants => {
			return { servants };
		});
	}
</script>

<script>
	import ServantTable from "../../components/ServantTable.svelte";
	import FilterBoard from "../../components/FilterBoard.svelte";

	function handleSelectors(event) {
		const selectors = event.detail;
		filteredServants = servants.filter(servant => {
			return selectors[servant.class.title] === true;
		})
	}
	export let servants;
	let filteredServants = servants;
</script>

<svelte:head>
	<title>Servants</title>
</svelte:head>

<h1>Servants</h1>

<FilterBoard servants={servants} on:select={handleSelectors} />
<ServantTable servants={filteredServants} />