<script lang="ts">
    import Window from "../layout/Window.svelte";
    import { Button, P } from "flowbite-svelte";
    import { ArrowRightOutline } from "flowbite-svelte-icons";
</script>

<Window heading="Home">
    <div class="flex-column">
        <P
            class="mb-6 text-lg text-white text-center lg:text-xl sm:px-16 xl:px-48 dark:text-gray-400"
            >This is the best Virtalization Manager you have ever tried, no
            doubt about it!</P
        >
        <Button href="#/settings">
            Let's get started
            <ArrowRightOutline class="w-6 h-6 ms-2" />
        </Button>
    </div>
</Window>
