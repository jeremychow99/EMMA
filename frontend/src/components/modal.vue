<template>
    <div class="modal" v-if="show">
        <div class="modal-content">
            <h3>New Equipment</h3>
                <form @submit.prevent="handleSubmit">
                    <label for="name">Name: </label>      
                    <input type="text" v-model="equipment_name">
                    <label for="location">Location:</label>
                    <input type="text" v-model="equipment_location">
                    <label for="status">Status:</label>
                    <input type="text" v-model="equipment_status">
                    <label for="last maintained">Last Maintained:</label>
                    <input type="text" v-model="last_maintained">
                    <v-btn type="submit" style="color: white; background-color: #5D5FEF">Add</v-btn>
                </form>
            </div>
    </div>
</template>

<script>
import axios from 'axios'
import { equipmentURL  } from '../../api'

export default{
    props: {
        show: Boolean
    },

    data() {
        return {
            equipment_name: "",
            equipment_location: "",
            equipment_status: "",
            last_maintained: "",
        }
    },

    methods: {
        handleSubmit(){
            var data = {
                equipment_name: this.equipment_name,
                equipment_location: this.equipment_location,
                equipment_status: this.equipment_status,
                last_maintained: this.last_maintained
            }
            axios.post(`${equipmentURL}`, data, {
                headers: {
                        "Content-Type": "application/json",
                        "Access-Control-Allow-Origin": "*",
                        Authorization: "Bearer " + localStorage.token,
                    }
                })
                .then(response => {
                    console.log(response.data)
                })
            this.equipment_name = ""
            this.equipment_location = ""
            this.equipment_status = ""
            this.last_maintained = ""
        }
    }
}

</script>