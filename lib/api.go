package lib

import (
	"fmt"
	"jetsend_opens/shared/helpers"
	"net/http"
	"time"

	"jetsend_opens/shared/log"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "open indexHandler")
}

//EmailOpen ...
func EmailOpen(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)
	///open/{email_id}/{account_id}/{url_index}/{url_digest}
	messageID := id["email_id"]
	accountID := id["account_id"]

	log.Info(r.Context(), "accID: ", accountID, " msgID: ", messageID, r.RemoteAddr, time.Now().String())

	writeToKafkaTopic(r.Context(), &helpers.Opens{AccountID: accountID, EmailUUID: messageID, Browser: r.UserAgent(), IPAddress: r.RemoteAddr, RecordedAT: time.Now().UTC()})
	http.Redirect(w, r, ServiceConfig.Service.PixelImage, http.StatusSeeOther)

}

//proto://domainname/click/email_id/url_index/url_digest
