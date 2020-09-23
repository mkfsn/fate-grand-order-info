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

	.star {
		color: #FFD500;
		text-shadow: 1px 1px 1px #000000;
		font-size: 110%;
		text-align: center;
	}
</style>

<svelte:head>
	<title>Servants</title>
</svelte:head>

<h1>Servants</h1>

<table>
    <thead>
		<tr>
			<th>編號</th>
			<th>圖示</th>
			<th>角色名稱</th>
			<th>職階</th>
			<th>星等</th>
			<th>基本ATK</th>
			<th>最大ATK</th>
			<th>基本HP</th>
			<th>最大HP</th>
		</tr>
	</thead>
	<tbody>
		{#each servants as servant}
		    <tr>
				<td>{servant.id}</td>
				<td><img src={servant.icon['src-1x']}  alt={servant.name.en} srcset="{servant.icon['src-1.5x']} 1.5x, {servant.icon['src-2x']} 2x"/></td>
				<td>{servant.name.jp}</td>
                <td><img src={servant.class['src-1x']} width="40" height="40" alt={servant.class.title} srcset="{servant.class['src-1.5x']} 1.5x, {servant.class['src-2x']} 2x"/></td>
				<td class="star">
					{#each Array(servant.rarity) as _}★{/each}
				</td>
				<td>{servant.attack.min}</td>
				<td>{servant.attack.max}</td>
				<td>{servant.hp.min}</td>
				<td>{servant.hp.max}</td>
			</tr>
		{/each}
	</tbody>
</table>