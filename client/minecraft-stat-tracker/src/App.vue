<template>
  <div>
    <h1>PlayerList</h1>
    <PlayerList :users="Users"/>
  </div>
</template>

<script>
import PlayerList from './components/PlayerList.vue'

export default {
  name: 'App',
  components: {
    PlayerList
    
  },
  data() {
        return {
            Users: []
        }
    },
    created() {
      fetch("http://localhost:8081/players")
        .then(response => response.json())
        .then(data => (this.Users = data));
        // this.Users = [
        //     {
        //         Username: "ShrubLord15",
        //         TotalTime: "99h99m99s",
        //         MessageCount: 999
        //     },
        //     {
        //         Username: "BDeltaC",
        //         TotalTime: "929h929m929s",
        //         MessageCount: 2131
        //     }
        // ]
    }
}
</script>

<style>
@import url("https://fontlibrary.org/face/minecraftia");

#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #ffffff;
  margin-top: 60px;
}

body {
  background-image: url("https://wallpaperaccess.com/full/6104111.png");
  font-family: Minecraftia;
}

</style>
