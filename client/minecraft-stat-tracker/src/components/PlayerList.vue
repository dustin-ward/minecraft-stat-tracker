<template>
    <div class="PlayerList">
        <table>
            <tr class="table-header">
                <th>Username</th>
                <th>Total Play Time</th>
                <th>Messages Sent</th>
            </tr>
            <PlayerListItem v-for="User in users" :key="User.Username" :user="User"/>
        </table>
    </div>
</template>

<script>
import PlayerListItem from './PlayerListItem.vue'
export default {
    name: 'PlayerList',
    components: { 
        PlayerListItem
    },
    props: {
        users: Array
    }
}
</script>

<style>
table {
  margin-left: auto;
  margin-right: auto;
}

.PlayerList {
  margin-left: auto;
  margin-right: auto;
  width: 50%;
  padding: 10px;
}

.table-header {
  background-color: #585858;
  color: rgb(243, 243, 243);
  text-align: center;
}

.table-header th {
  padding: 0px 10px;
}
</style>