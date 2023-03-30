<template>
    <navbar></navbar>
    <v-container>
        <v-row>
            <v-col cols="4">
                <h2 class="text-start mt-3">View</h2>
                <h2 class="text-start mb-3">Maintenance</h2>
            </v-col>
            <v-col cols="8">
            </v-col>
        </v-row>
        
    </v-container>
    
</template>

<script>
import navbar from "../components/navbar.vue"
import { equipmentURL  } from '../../api'
import axios from "axios";

export default {
    components: {
        navbar
    },

    data(){ 
        return {
            eq_result: [],

        }
    },

    async mounted() {
        let eq_result = await axios.get(equipmentURL)
        this.eq_result = eq_result
        console.log(this.eq_result)
    }

}

</script>