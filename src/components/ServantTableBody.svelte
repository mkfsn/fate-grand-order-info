<script>
    import Image from "./Image.svelte";
    export let servants;
</script>

<style>
    @media (max-width: 576px) {
        .hide-xs {
            display: none;
        }
    }

    .class > :global(img) {
        width: 5vw;
        height: 5vw;
        max-height: 2em;
        max-width: 2em;
    }

    .icon > :global(img) {
        width: 10vw;
        height: 10vw;
        max-height: 3em;
        max-width: 3em;
    }

    .star {
        color: #FFD500;
        text-shadow: 1px 1px 1px #000000;
        font-size: 110%;
        text-align: center;
    }

    .hidden {
        display: none;
    }

    .commands {
        min-width: min(21vw, 6em);
        text-align: center;
    }
    .commands > :global(img) {
        margin-left: -3.5em;
        width: 7vw;
        height: 7vw;
        max-height: 2.5em;
        max-width: 2.5em;
    }
    .commands > :global(img:first-child) {
        margin-left: 0;
    }
    .nobel-phantasm > :global(img) {
        width: 7vw;
        height: 7vw;
        max-height: 2.5em;
        max-width: 2.5em;
    }
</style>

<tbody>
{#each servants as servant (servant.id) }
    <tr class="{servant.hidden === true ? 'hidden': ''}">
        <td class="hide-xs">{servant.id}</td>
        <td class="icon">
            <Image title={servant.name.en} sources={servant.icon}/>
        </td>
        <td>{servant.name.jp}</td>
        <td class="class">
            <Image title={servant.class.title} sources={servant.class}/>
        </td>
        <td class="star">
            {#each Array(servant.rarity) as _}★{/each}
        </td>
        <td class="hide-xs">{servant.attack.min}</td>
        <td>{servant.attack.max}</td>
        <td class="hide-xs">{servant.hp.min}</td>
        <td>{servant.hp.max}</td>
        <td class="commands">
            {#each servant.command_cards as command}
                <Image sources={command} />
            {/each}
        </td>
        <td class="nobel-phantasm">
            <Image title={servant.noble_phantasm.title} sources={servant.noble_phantasm}/>
        </td>
    </tr>
{/each}
</tbody>
