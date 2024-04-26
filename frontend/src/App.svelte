<script lang="ts">
    import { Greet } from "../wailsjs/go/app/App.js";
    import Sidebar from "./components/sidebar.svelte";

    let resultText: string = "Please enter your name below 👇";
    let name: string;

    function greet(): void {
        Greet(name).then((result) => (resultText = result));
    }
</script>

<main class="h-full">
    <Sidebar />
</main>
