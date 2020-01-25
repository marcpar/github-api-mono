package main

import (
	"net/http"
	"testing"
)


// TEST: 1. Creating a Github organization
func TestCreateOrganization(t *testing.T) {

	req, _ := http.NewRequest("POST", "/organization", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}

}

// TEST: 2. Persisting Github comments against a given Github organization
func TestCreateComments(t *testing.T){

}

// TEST: 3. Returning an array of all the comments that have been registered against a Github organization
func TestListComments(t *testing.T){

}

// TEST: 4. Soft-deleting all comments associated with a particular organization. We define a "soft delete" to mean that deleted items should not be returned in GET calls, but should remain in the database for emergency retrieval and audit purposes. 
func TestDeleteComment(t *testing.T){

}

// TEST: 5. Returning an array of members of an organization (with their login, avatar URL, the numbers of followers they have, and the number of people they're following), sorted in descending order by the number of followers. 
func TestMemberInfo(t *testing.T){

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}




