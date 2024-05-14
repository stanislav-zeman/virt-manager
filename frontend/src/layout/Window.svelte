<script lang="ts">
    import { Heading } from "flowbite-svelte";

    export let heading: string;
</script>

<div class="w-4/5">
    <Heading tag="h2" class="m-6 text-white">{heading}</Heading>
    <slot />
</div>
