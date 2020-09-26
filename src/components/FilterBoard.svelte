<script>
    import { selectedClasses } from './store.js';
    export let servants;

    const basis = [
        "saber", "archer", "lancer", "rider", "caster", "assassin", 'berserker',
        'shielder', 'ruler', 'avenger', 'alterego', 'mooncancer', 'foreigner', 'beast',
    ];

    function sortByClass(a, b) {
        const normalize = (s) => s.toLowerCase().replace(/-/g, '');
        let x = basis.indexOf(normalize(a)),
            y = basis.indexOf(normalize(b));
        x = x === -1 ? 99 : x;
        y = y === -1 ? 99 : y;
        return x < y ? -1 : 1;
    }

    let orderedClasses = [...new Set(servants.map(s => s.class.title)).keys()].sort(sortByClass);
    orderedClasses.forEach(v => {
        selectedClasses.update(s => Object.assign({}, s, {[v]: true}));
    })

    function toggleClass(event) {
        selectedClasses.update(s => Object.assign({}, s, {[event]: !s[event]}));
    }
</script>

<style>
    div {
        margin-top: 1em;
        margin-bottom: 1em;
    }
    img {
        opacity: 0.3;
        width: 6%;
        height: 6%;
        max-width: 3em;
        max-height: 3em;
    }
    img:hover {
        cursor: pointer;
    }
    img.selected {
        opacity: 1;
    }
</style>

<div>
    {#each Object.entries($selectedClasses) as entry}
        <img
            class="{entry[1] ? 'selected' : ''}"
            on:click={toggleClass(entry[0])}
            src="/images/Icon_Class_{entry[0]}_Gold.png"
            alt={entry[0]}
        />
    {/each}
</div>