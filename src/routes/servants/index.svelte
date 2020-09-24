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
		width: 180px;
		text-align: center;
	}
	.commands img {
		height: 60px;
		width: 60px;
		margin-left: -30px;
	}
	.commands img:first-child {
		margin-left: 0;
	}
	.nobel-phantasm img {
		height: 60px;
		width: 60px;
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
				<td><img src={servant.icon['src-1x']}  alt={servant.name.en} srcset="{servant.icon['src-1.5x']} 1.5x, {servant.icon['src-2x']} 2x"/></td>
				<td>{servant.name.jp}</td>
                <td>
					<img src={servant.class['src-1x']} width="40" height="40" alt={servant.class.title} srcset="{servant.class['src-1.5x']} 1.5x, {servant.class['src-2x']} 2x"/>
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
						<img src={command['src-1x']} width="40" height="40" alt={command.title} srcset="{command['src-1.5x']} 1.5x, {command['src-2x']} 2x"/>
					{/each}
				</td>
				<td class="nobel-phantasm">
					<img src={servant.noble_phantasm['src-1x']} width="40" height="40" alt={servant.noble_phantasm.title} srcset="{servant.noble_phantasm['src-1.5x']} 1.5x, {servant.noble_phantasm['src-2x']} 2x"/>
				</td>
			</tr>
		{/each}
	</tbody>
</table>