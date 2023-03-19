<template>
    <navbar></navbar>

    <v-container>

        <v-row>
            <v-col cols="4">
                <h2 class="text-start mt-3">Schedule</h2>
                <h2 class="text-start mb-3">Maintenance</h2>
            </v-col>
            <v-col cols="8">
                <v-row class="mt-0 mx-2">
                    <v-col cols="5" style="padding:8px 16px">
                        <p class="text-right font-weight-regular">EquipmentID: </p>
                    </v-col>
                    <v-col cols="7" style="padding:8px 16px">
                        <p>1234</p>
                    </v-col>
                </v-row>
                <v-row class="mt-0 mx-2">
                    <v-col cols="5" style="padding:8px 16px">
                        <p class="text-right font-weight-regular">Parts: </p>
                    </v-col>
                    <v-col cols="7" style="padding:8px 16px">
                        <p>Reserved</p>
                    </v-col>
                </v-row>
                <v-row class="mt-0 mx-2">
                    <v-col cols="5" style="padding:8px 16px">
                        <p class="text-right font-weight-regular">Maintenance Datetime: </p>
                    </v-col>
                    <v-col cols="7" style="padding:8px 16px">
                    </v-col>
                </v-row>
            </v-col>
        </v-row>
        

        <v-tabs class="mx-auto mt-3" align-tabs="center" v-model="tab">
            <v-tab value="equipment">
                Equipment
            </v-tab>
            <v-tab value="parts">
                Parts
            </v-tab>
        </v-tabs>

        <v-window v-model="tab">

            <v-window-item key="equipment" value="equipment">
                <v-container class="rounded-lg" style="background-color: lightblue; padding: 5px">
                
                    <v-text-field
                        :loading="loading"
                        density="compact"
                        variant="solo"
                        label="Search equipment name"
                        single-line
                        hide-details
                        @click:append-inner="onClick"
                        class="my-2"
                        style="margin:auto; width: 50%;"
                    ></v-text-field>
                   
                        

                    <v-table fixed-header height="400px">
                        <thead>
                            <tr>
                                <th class="text-center">
                                    Equipment ID
                                </th>
                                <th class="text-center">
                                    Equipment Name
                                </th>
                                <th class="text-center">
                                    
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr
                                v-for="eqp in eqp_list"
                                :key="eqp.id"
                            >
                                <td>{{ item.name }}</td>
                                <td>{{ item.calories }}</td>
                            </tr>
                        </tbody>
                    </v-table>
                </v-container>
            </v-window-item>

            <v-window-item key="parts" value="parts">
                <v-container class="rounded-lg" style="background-color: lightgray; padding: 5px">
                
                <v-text-field
                    :loading="loading"
                    density="compact"
                    variant="solo"
                    label="Search part name"
                    single-line
                    hide-details
                    @click:append-inner="onClick"
                    class="my-2"
                    style="margin:auto; width: 50%;"
                ></v-text-field>
               
                    

                <v-table fixed-header height="400px">
                    <thead>
                        <tr>
                            <th class="text-center">
                                Part ID
                            </th>
                            <th class="text-center">
                                Part Name
                            </th>
                            <th class="text-center">
                                Quantity
                            </th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr
                            v-for="eqp in eqp_list"
                            :key="eqp.id"
                        >
                            <td>{{ item.name }}</td>
                            <td>{{ item.calories }}</td>
                        </tr>
                    </tbody>
                </v-table>
            </v-container>
            </v-window-item>
        </v-window>
    </v-container>


    <link rel="stylesheet" href="../../src/stylesheets/schedule_style.css">
</template>

<script>
import navbar from "../components/navbar.vue"

export default {
    components: {
        navbar
    },

    data() {
        return {
            tab: null,
            date: null
        }
    }
};
</script>