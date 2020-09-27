<script>
    import {tweened} from 'svelte/motion';
    import ServantTableHead from "./ServantTableHead.svelte";
    import ServantTableBody from "./ServantTableBody.svelte";

    export let servants;
    const sizeof = (arr) => arr ? arr.filter(v => !v.hidden).length : 0;
    const amount = tweened(sizeof(servants), {
        duration: 250
    });
    $: amount.set(sizeof(servants));
</script>

<style>
    .container {
        overflow-x: auto;
    }

    table {
        margin: 0 0 1em 0;
        line-height: 1.5;
        border-collapse: collapse;
        table-layout: auto;
        font-variant-numeric: lining-nums tabular-nums;
    }

    @media (max-width: 576px) {
        table {
            font-size: .7em;
        }
    }

    table :global(td),
    table :global(th) {
        border: 1px solid #111111;
        padding-left: .5em;
        padding-right: .5em;
    }
</style>

<div class="container">
    <table>
        <caption>共 { Math.floor($amount) } 筆資料</caption>
        <ServantTableHead/>
        <ServantTableBody servants={servants}/>
    </table>
</div>
