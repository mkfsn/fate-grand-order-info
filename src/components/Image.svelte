<script>
    export let title = "";
    export let sources = {
        'src-1x': '',
        'src-1.5x': '',
        'src-2x': ''
    };
    export let width = 40;
    export let height = 40;
    $: srcset = `${sources['src-1.5x']} 1.5x, ${sources['src-2x']} 2x`;
</script>

<img src={sources['src-1x']} width={width} height={height} alt={title} srcset={srcset}/>
