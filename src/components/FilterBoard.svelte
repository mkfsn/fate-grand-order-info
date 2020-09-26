<script>
    import {createEventDispatcher} from 'svelte';

    export let servants;
    let orderedClasses;
    let selectedClasses = {};
    const basis = [
        "saber", "archer", "lancer", "rider", "caster", "assassin", 'berserker',
        'shielder', 'ruler', 'avenger', 'alterego', 'mooncancer', 'foreigner', 'beast',
    ];
    $: {
        orderedClasses = [...new Set(servants.map(servant => servant.class.title)).keys()];
        orderedClasses.sort((a, b) => {
            const normalize = (s) => {
                return s.toLowerCase().replace(/-/g, '');
            };
            let x = basis.indexOf(normalize(a)), y = basis.indexOf(normalize(b));
            if (x === -1) {
                x = 99;
            }
            if (y === -1) {
                y = 99;
            }
            if (x < y) {
                return -1;
            }
            return 1;
        });

        selectedClasses = {}
        orderedClasses.forEach((v) => {
            selectedClasses[v] = true;
        });
    }

    const dispatch = createEventDispatcher();
    function toggleClass(event) {
        selectedClasses[event] = !selectedClasses[event];
        dispatch('select', selectedClasses);
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
    {#each orderedClasses as _class}
        <img
            class="{selectedClasses[_class] ? 'selected' : ''}"
            on:click={toggleClass(_class)}
            src="/images/Icon_Class_{_class}_Gold.png"
            alt={_class}
        />
    {/each}
</div>