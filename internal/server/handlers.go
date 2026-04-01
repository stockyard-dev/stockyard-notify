package server
import("encoding/json";"net/http";"strconv";"github.com/stockyard-dev/stockyard-notify/internal/store")
func(s *Server)handleListTemplates(w http.ResponseWriter,r *http.Request){list,_:=s.db.ListTemplates();if list==nil{list=[]store.Template{}};writeJSON(w,200,list)}
func(s *Server)handleCreateTemplate(w http.ResponseWriter,r *http.Request){var t store.Template;json.NewDecoder(r.Body).Decode(&t);if t.Name==""||t.Body==""{writeError(w,400,"name and body required");return};if t.Channel==""{t.Channel="sms"};s.db.CreateTemplate(&t);writeJSON(w,201,t)}
func(s *Server)handleSend(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);var req struct{Recipient string `json:"recipient"`};json.NewDecoder(r.Body).Decode(&req);if req.Recipient==""{writeError(w,400,"recipient required");return};send:=&store.Send{TemplateID:id,Recipient:req.Recipient,Status:"sent"};s.db.RecordSend(send);writeJSON(w,200,send)}
func(s *Server)handleGetSends(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);list,_:=s.db.ListSends(id);if list==nil{list=[]store.Send{}};writeJSON(w,200,list)}
func(s *Server)handleDelete(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.DeleteTemplate(id);writeJSON(w,200,map[string]string{"status":"deleted"})}
func(s *Server)handleOverview(w http.ResponseWriter,r *http.Request){m,_:=s.db.Stats();writeJSON(w,200,m)}
