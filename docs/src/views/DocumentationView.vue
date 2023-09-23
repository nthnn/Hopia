<template>
    <h1>Command-Line Interface</h1>
    <hr/>

    <div class="card card-body bg-dark rounded-0 border-white">
        <h3>DDoS Attack</h3>
        <p>The <code>ddos</code> command allows you to launch a Distributed Denial of Service (DDoS) attack against a target URL using a specified configuration. DDoS attacks aim to overwhelm a target server with a flood of traffic, making it temporarily or permanently unavailable.</p>
        <br/>

        <b>Usage:</b>
        <code>hopia ddos &lt;url&gt; &lt;config&gt;</code>
        <br/>

        <b>Arguments:</b>
        <ul>
            <li>
                <code>&lt;url&gt;</code>
                <p>The target URL of the system you want to test. Ensure you have proper authorization to test this system.</p>
            </li>
            <li>
                <code>&lt;config&gt;</code>
                <p>The configuration file specifying the parameters of the DDoS attack. This file should define attack characteristics such as duration, intensity, and attack vectors. The <code>&lt;config&gt;</code> must be an existing *.properties file that contains the following keys below:</p>
                <br/>

                <ul>
                    <li><code>Thread-Count</code> - Number of thread counts that will ping the target.</li>
                    <li><code>*</code> - Other than <code>Thread-Code</code>, other keys will be used as HTTP request header.</li>
                </ul>
            </li>
        </ul>

        <b>Example:</b>
        <code>./hopia ddos http://localhost attack.properties</code>
    </div>
    <br/>

    <div class="card card-body bg-dark rounded-0 border-white">
        <h3>Hash Crack</h3>
        <p>The <code>crack</code> command allows you to attempt to brute force an encrypted text (hash) using a text list. This feature is often used for password recovery or hash analysis, but it should only be used for legitimate purposes on systems where you have explicit authorization.</p>
        <br/>

        <b>Usage:</b>
        <code>hopia crack &lt;hash&gt; &lt;config&gt;</code>
        <br/>

        <b>Arguments:</b>
        <ul>
            <li>
                <code>&lt;hash&gt;</code>
                <p>The encrypted text (hash) you want to crack.</p>
            </li>
            <li>
                <code>&lt;config&gt;</code>
                <p>The configuration file specifying the parameters of the brute force attack. This file should define settings such as the hashing algorithm and the text list for attempting to crack the hash. The &lt;config&gt; must be an existing *.properties file that contains the following keys below:</p>
                <ul>
                    <li><code>algorithm</code> - The algorithm of hash to be used. Possible values are: <code>md5</code>, <code>sha1</code>, and <code>sha256</code>.</li>
                    <li><code>dictionary</code> - Dictionary file to be used on brute forcing.</li>
                </ul>
            </li>
        </ul>

        <b>Example:</b>
        <code>./hopia crack 596a96cc7bf9108cd896f33c44aedc8a bruteforce.properties</code>
    </div>
</template>

<style>
    .table-docs {
        background-color: var(--bs-dark) !important;
        color: var(--bs-white) !important;
    }
</style>