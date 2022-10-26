FROM vectorized/redpanda:v22.1.9

USER nobody

CMD ["redpanda", "start", "--overprovisioned", "--smp", "1", "--memory", "1G", "--reserve-memory", "0M", "--node-id", "0", "--check=false"]
