<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  list-style: none;
}

a {
  text-decoration: none;
  color: inherit;
}

#app {
  width: 100%;
  height: 100%;
}

::-webkit-scrollbar {
  display: none;
}
</style>
