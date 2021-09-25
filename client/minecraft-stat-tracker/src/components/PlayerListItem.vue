<template>
    <tr class="table-row">
        <td class="username">
          <a href=""><img :src="`http://localhost:8081/static/images/${user.Username}_face.png`"
                          width="25" height="25"
                          :alt="user.Username" class="pfp">{{user.Username}}</a>
        </td>
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

.username {
  text-align: left;
}
</style>