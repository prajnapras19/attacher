package attachment

type Attachment struct {
	ID       uint
	Serial   string
	Path     string
	Name     string
	IsActive bool
	UserID   uint
}

func (a *Attachment) ToAttachmentResponse() AttachmentResponse {
	return AttachmentResponse{
		Name: a.Name,
		Link: "./attachments/" + a.Serial,
		// TODO: support external link in the future
		// the attachment type should be added to differentiate local and external file
	}
}

type AttachmentResponse struct {
	Name string
	Link string
}
