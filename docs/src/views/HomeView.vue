<script setup lang="ts">
import { ArrowRightIcon } from '@heroicons/vue/24/outline';
</script>

<template>
    <div class="row">
        <div class="col-lg-6">
            <br class="desktop-only" />
            <h1>Hopia</h1>
            <hr/>

            <p>Hopia is a powerful and versatile penetration testing tool designed for ethical and legitimate security testing purposes. It provides a range of security testing capabilities to help assess and enhance the security of systems and networks.</p>
            <p>It is also important to note that Hopia is intended to be used responsibly, within the boundaries of the law, and with proper authorization from system owners or administrators. Any malicious or unauthorized use of this software is strictly prohibited and may result in legal consequences.</p>

            <RouterLink to="/docs" class="btn btn-outline-primary learn-more-btn">
                Learn More
                <ArrowRightIcon class="icon-18" />
            </RouterLink>
        </div>

        <div class="col-lg-6">
            <br class="desktop-only" /><br/>

            <center>
                <img src="@/assets/hopia.png" :width="280" alt="Hopia Logo" />
            </center>
        </div>
    </div>
</template>

<style scoped>
    .desktop-only {
        display: block;
    }

    @media only screen and (max-width: 600px) {
        .desktop-only {
            display: none;
        }

        .learn-more-btn {
            width: 100%;
        }
    }
</style>