<script>
    export let title = "";
    export let sources = {
        'src-1x': '',
        'src-1.5x': '',
        'src-2x': ''
    };
    $: srcset = `${sources['src-1.5x']} 1.5x, ${sources['src-2x']} 2x`;
</script>

<img src={sources['src-1x']} alt={title} srcset={srcset}/>
