<template>
  <div :style="initial" jid="app">
    <FileParser />
  </div>
</template>

<script>
import FileParser from "./components/FileParser.vue";
import "./assets/css/main.css";

export default {
  name: "app",
  data() {
    return {
      initial: {
        backgroundColor: "black"
      }
    };
  },
  components: {
    FileParser
  }
};
</script>
<style scoped>
body {
  background-color: red;
}
</style>
