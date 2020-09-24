<script context="module">
	export function preload({ params, query }) {
		return this.fetch(`/data/servants.json`).then(r => r.json()).then(servants => {
			return { servants };
		});
	}
</script>

<script>
	import Image from "../../components/Image.svelte";
	export let servants;
</script>

<style>
	table {
		margin: 0 0 1em 0;
		line-height: 1.5;
		border-collapse: collapse;
	}

	table td, table th {
		border: 1px solid #111111;
		padding-left: .5em;
		padding-right: .5em;
	}

	.star {
		color: #FFD500;
		text-shadow: 1px 1px 1px #000000;
		font-size: 110%;
		text-align: center;
	}

	.commands {
		width: 120px;
		text-align: center;
	}
	.commands > :global(img) {
		margin-left: -20px;
	}
	.commands > :global(img:first-child) {
		margin-left: 0;
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
			<th>基本 ATK</th>
			<th>最大 ATK</th>
			<th>基本 HP</th>
			<th>最大 HP</th>
			<th>指令卡</th>
			<th>寶具</th>
		</tr>
	</thead>
	<tbody>
		{#each servants as servant}
		    <tr>
				<td>{servant.id}</td>
				<td>
					<Image width=70 height=70 title={servant.name.en} sources={servant.icon}/>
				</td>
				<td>{servant.name.jp}</td>
				<td>
					<Image width=30 height=30 title={servant.class.title} sources={servant.class}/>
				</td>
				<td class="star">
					{#each Array(servant.rarity) as _}★{/each}
				</td>
				<td>{servant.attack.min}</td>
				<td>{servant.attack.max}</td>
				<td>{servant.hp.min}</td>
				<td>{servant.hp.max}</td>
				<td class="commands">
					{#each servant.command_cards as command}
						<Image width=40 height=40 sources={command} />
					{/each}
				</td>
				<td class="nobel-phantasm">
					<Image width=40 height=40 title={servant.noble_phantasm.title} sources={servant.noble_phantasm}/>
				</td>
			</tr>
		{/each}
	</tbody>
</table>