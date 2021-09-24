<template>
    <tr class="table-row">
        <td><a href=""><img :src="user.FacePic" :alt="user.Username" class="pfp"></a>{{user.Username}}</td>
        <td>{{user.TotalTime}}</td>
        <td>{{user.MessageCount}}</td>
    </tr>
</template>

<script>
export default {
    name: 'PlayerListItem',
    props: {
        user: Object
    },
}
</script>

<style>
.table-row td {
  background-color: #c2c2c2;
  color: #0e0e0e;
  padding: 12px 15px;
}

.table-row a {
  color: rgb(255, 123, 0);
}

.pfp {
  padding: 0px 10px;
}

.main-pfp {
  display: block;
  margin-left: auto;
  margin-right: auto;
}
</style>