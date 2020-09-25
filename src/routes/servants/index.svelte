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

	export let servants;
	let selectors = undefined;
	function handleSelectors(event) {
		selectors = event.detail;
	}
	$: filteredServants = (selectors === undefined) ? servants : servants.map(servant => {
		servant.hidden = selectors[servant.class.title] !== true;
		return servant;
	})
</script>

<svelte:head>
	<title>Servants</title>
</svelte:head>

<h1>Servants</h1>

<FilterBoard servants={servants} on:select={handleSelectors} />
<ServantTable servants={filteredServants} />