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
        this.Users = [
            {
                Username: "ShrubLord15",
                TotalTime: "99h99m99s",
                MessageCount: 999,
                FacePic: "https://randomuser.me/api/portraits/men/10.jpg"
            },
            {
                Username: "BDeltaC",
                TotalTime: "929h929m929s",
                MessageCount: 2131,
                FacePic: "https://randomuser.me/api/portraits/men/10.jpg"
            }
        ]
    }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
