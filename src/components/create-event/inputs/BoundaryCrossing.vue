<template>
    <div>
        <TeamSelection :model="model" label="By: "/>
        <b-button variant="primary"
                  @click="sendEvent()"
                  :disabled="model.team === null">
            Add
        </b-button>
    </div>
</template>

<script>
    import TeamSelection from "@/components/common/TeamSelection";
    import {submitGameEvent} from "../../../submit";

    export default {
        name: "BoundaryCrossing",
        components: {TeamSelection},
        data() {
            return {
                model: {
                    team: null,
                }
            }
        },
        methods: {
            sendEvent: function () {
                submitGameEvent({
                    type: 'BOUNDARY_CROSSING',
                    boundaryCrossing: {
                        by_team: this.model.team.toLocaleUpperCase(),
                    }
                });
                this.$root.$emit('bv::hide::modal', 'new-event-modal');
            }
        },
    }
</script>

<style scoped>

</style>
