<script setup lang="ts">
import { ArrowRightIcon } from '@heroicons/vue/24/outline';
import HomePageCard from "../components/HomePageCard.vue";
</script>

<template>
    <div class="row">
        <div class="col-lg-7">
            <br class="desktop-only" />
            <h1>Hopia</h1>
            <hr/>

            <p>Hopia is a powerful and versatile wannabe penetration testing tool designed for ethical and legitimate security testing purposes. It provides a range of wannabe security testing capabilities to help assess and enhance the security of systems and networks.</p>
            <p>It is also important to note that Hopia is intended to be used responsibly, within the boundaries of the law, and with proper authorization from system owners or administrators. Any malicious or unauthorized use of this software is strictly prohibited and may result in legal consequences.</p>

            <RouterLink to="/docs" class="btn btn-outline-light border-1 learn-more-btn">
                Learn More
                <ArrowRightIcon class="icon-18" />
            </RouterLink>
        </div>

        <div class="col-lg-5">
            <br/>

            <center>
                <img src="@/assets/hopia.png" :width="270" alt="Hopia Logo" />
                <br/>
                <small>Hopia is a popular Philippine bean-filled moon cake-like pastry.</small>
            </center>
        </div>
    </div>
    <br/><br/><br/>

    <h1>Features</h1>
    <hr/>
    <div class="row">
        <HomePageCard
            title="DDoS Testing"
            content="Simulate Distributed Denial of Service (DDoS) attacks to assess the resilience of your systems against such threats. Only use this feature with proper authorization on systems you own or have explicit permission to test." />

        <HomePageCard
            title="Hash Cracking"
            content="Attempt to crack hash strings using popular algorithms like MD5, SHA-1, and SHA-256. This feature should only be used for legitimate purposes, such as password recovery on systems where you have explicit authorization." />

        <HomePageCard
            title="Admin Page Finder"
            content="Find and discover administrator login pages on websites using a file containing list of common URLs. Always use this feature responsibly and only on websites for which you have proper and legal authorization." />

        <HomePageCard
            title="Port Scanning"
            content="Scan for open ports from a given host URL in order to identify potential vulnerabilities. As always, ensure you have legal permission and authorization to scan the target system before using this feature." />
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