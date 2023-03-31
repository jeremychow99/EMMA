<template>
    <navbar></navbar>
    
    <v-container>
      <div class="text-h4 text-center font-weight-black my-5">Equipment</div>
    </v-container>
    
    <v-container>
        <v-row>
            <v-col class="text-h6" cols="4">
                All Equipment Overview
            </v-col>
            <v-col cols="5">
            </v-col>
            <v-col style="margin-bottom: 10px;">
                <v-btn @click="show_modal = true" style="background-color: #5D5FEF; color: white;">
                    Add Equipment
                </v-btn>
                <modal :show="show_modal" @add_equipment="add_equipment"></modal>
            </v-col>
        </v-row>

            <v-table fixed-header height="300px">
                <thead>
                    <tr>
                        <th>
                            Equipment Name
                        </th>
                        <th>
                            Equipment Location
                        </th>                        
                        <th>
                            Last Maintained
                        </th>                        
                        <th>
                            Status
                        </th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="equipment of eq_arr">
                        <td>
                            {{ equipment.equipment_name }}
                        </td>
                        <td>
                            {{ equipment.equipment_location }}
                        </td>
                        <td>
                            {{ equipment.last_maintained }}
                        </td>
                        <td>
                            {{ equipment.equipment_status }}
                        </td>
                    </tr>
                </tbody>
            </v-table>
    </v-container>
    
</template>

<script>
import navbar from "../components/navbar.vue"
import { equipmentURL  } from '../../api'
import axios from "axios";
import modal from "../components/modal.vue"

export default {
    components: {
        navbar,
        modal
    },

    data(){ 
        return {
            eq_arr: [],
            show_modal: false,
            new_equipment: {
                equipment_name: "",
                equipment_location: "",
                equipment_status: "",
                last_maintained: ""
            }
        }
    },

    methods: {
        add_equipment(equipment){
            console.log("=== START add_equipment() ===")


        }, 
        showModalForm() {
            this.show_modal = true;
        },
    },

    async mounted() {
        let eqp_result = await axios.get(equipmentURL)
        this.eq_arr = eqp_result.data.data.equipment
        console.log(this.eq_arr)
    }

}

</script>