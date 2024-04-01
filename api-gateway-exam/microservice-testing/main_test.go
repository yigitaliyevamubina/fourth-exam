package microservicetesting

import (
	microservicetesting "exam/api-gateway/microservice-testing/microservice-testing"
	"testing"
)

func TestRunApiTest(t *testing.T) {
	microservicetesting.RunApiTest(t)
}