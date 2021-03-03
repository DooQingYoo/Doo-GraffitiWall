<template>
  <div class="big-daddy">
    <Menu></Menu>
    <router-view/>
  </div>
</template>
<script>
import Menu from "@/components/Menu";
export default {
  components: {Menu},
}
</script>
<style>
.big-daddy {
  max-width: 1200px;
  margin: 0 auto;
}
</style>
