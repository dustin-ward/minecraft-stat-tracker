<template>
    <div class="PlayerList">
        <table>
            <tr>
                <td>Username</td>
                <td>Total Play Time</td>
                <td>Messages Sent</td>
            </tr>
            <tr :key="User.Username" v-for="User in users">
                <!-- <p>{{User.Username}}</p> -->
                <PlayerListItem :user="User"/>
            </tr>
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

</style>