package main

const HTML = `
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/water.css@2/out/water.css">

<main>
    <h1>Privacy-friendly statistical analysis ;)</h1>

    <h2>Your Weight</h2>
    <form action="/" method="post">
        <label for="name">Name</label>
        <input type="text" name="name" id="name">

        <label for="password">Password</label>
        <input type="password" name="password" id="password">

        <label for="weight">Weight</label>
        <input type="number" name="weight" id="weight">


        <input type="submit" value="submit">
    </form>

    <h2>Average Wieght</h2>
    <p id="output">%v</p>
</main>

<script>
</script>`
