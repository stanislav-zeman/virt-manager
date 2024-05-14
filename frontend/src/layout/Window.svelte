<script lang="ts">
    import { Heading } from "flowbite-svelte";

    export let heading: string;
</script>

<div class="w-4/5">
    <Heading tag="h3" class="my-6 text-white">{heading}</Heading>
    <hr class="w-3/5 border-gray-200 sm:mx-auto dark:border-gray-700 lg:my-6" />
    <slot />
</div>
