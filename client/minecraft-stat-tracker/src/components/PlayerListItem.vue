<template>
    <div class="PlayerListItem">
        <h2>{{user.Username}}</h2>
        <!-- <tr>
            <td><a href=""><img :src="user.FacePic" :alt="user.Username"></a>{{user.Username}}</td>
            <td>{{user.TotalTime}}</td>
            <td>{{user.MessageCount}}</td>
        </tr> -->
    </div>
</template>

<script>
export default {
    name: 'PlayerListItem',
    props: {
        user: Object
    }
}
</script>

<style>

</style>