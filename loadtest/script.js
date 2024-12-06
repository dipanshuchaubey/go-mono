import http from "k6/http"
import { check, sleep } from "k6"
import { Rate } from "k6/metrics"

// Custom metric to track error rate
export let errorRate = new Rate("errors")

// Define load stages
export let options = {
  stages: [
    { duration: "15s", target: 1000 }, // Ramp-up to 10 users over 30 seconds
    { duration: "15s", target: 1100 }, // Stay at 10 users for 1 minute
    { duration: "15s", target: 0 } // Ramp-up to 20 users over 30 seconds
  ],
  thresholds: {
    http_req_duration: ["p(95)<500"], // 95% of requests should be below 500ms
    errors: ["rate<0.01"] // Error rate should be less than 1%
  }
}

export default function () {
  const url = "http://localhost:56719/get"
  const response = http.get(url)

  // Check if the response status is 200
  const isSuccess = check(response, {
    "status is 200": r => r.status === 200,
    "response time < 500ms": r => r.timings.duration < 500
  })

  // Record errors if the request was unsuccessful
  errorRate.add(!isSuccess)

  // Pause between requests to simulate real user behavior
  sleep(1)
}
