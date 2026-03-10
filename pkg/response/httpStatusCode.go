package response

const (
	// Continue - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.2.1
	// This interim response indicates that everything so far is OK and that the client should continue with the request or ignore it if it is already finished.
	StatusContinue = 100

	// SwitchingProtocols - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.2.2
	// This code is sent in response to an Upgrade request header by the client, and indicates the protocol the server is switching too.
	StatusSwitchingProtocols = 101

	// Processing - Official Documentation @ https://tools.ietf.org/html/rfc2518#section-10.1
	// This code indicates that the server has received and is processing the request, but no response is available yet.
	StatusProcessing = 102

	// OK - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.3.1
	// The request has succeeded.
	StatusOK = 200

	// Created - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.3.2
	// The request has succeeded and a new resource has been created as a result of it.
	StatusCreated = 201

	// Accepted - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.3.3
	// The request has been received but not yet acted upon.
	StatusAccepted = 202

	// NonAuthoritativeInformation - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.3.4
	// This response code means returned meta-information set is not exact set as available from the origin server.
	StatusNonAuthoritativeInformation = 203

	// NoContent - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.3.5
	// There is no content to send for this request, but the headers may be useful.
	StatusNoContent = 204

	// ResetContent - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.3.6
	// This response code is sent after accomplishing request to tell user agent reset document view.
	StatusResetContent = 205

	// PartialContent - Official Documentation @ https://tools.ietf.org/html/rfc7233#section-4.1
	// This response code is used because of range header sent by the client to separate download into multiple streams.
	StatusPartialContent = 206

	// MultiStatus - Official Documentation @ https://tools.ietf.org/html/rfc2518#section-10.2
	// A Multi-Status response conveys information about multiple resources in situations where multiple status codes might be appropriate.
	StatusMultiStatus = 207

	// MultipleChoices - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.4.1
	// The request has more than one possible responses.
	StatusMultipleChoices = 300

	// MovedPermanently - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.4.2
	// This response code means that URI of requested resource has been changed permanently.
	StatusMovedPermanently = 301

	// MovedTemporarily - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.4.3
	// This response code means that URI of requested resource has been changed temporarily.
	StatusMovedTemporarily = 302

	// SeeOther - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.4.4
	// Server sent this response to directing client to get requested resource to another URI with a GET request.
	StatusSeeOther = 303

	// NotModified - Official Documentation @ https://tools.ietf.org/html/rfc7232#section-4.1
	// This is used for caching purposes. It tells the client that the response has not been modified.
	StatusNotModified = 304

	// UseProxy - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.4.6
	// Deprecated: Was defined to indicate that a requested response must be accessed by a proxy.
	StatusUseProxy = 305

	// TemporaryRedirect - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.4.7
	// Server sent this response to directing client to get requested resource to another URI with same method used prior request.
	StatusTemporaryRedirect = 307

	// PermanentRedirect - Official Documentation @ https://tools.ietf.org/html/rfc7538#section-3
	// This means that the resource is now permanently located at another URI.
	StatusPermanentRedirect = 308

	// BadRequest - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.5.1
	// This response means that server could not understand the request due to invalid syntax.
	StatusBadRequest = 400

	// Unauthorized - Official Documentation @ https://tools.ietf.org/html/rfc7235#section-3.1
	// The client must authenticate itself to get the requested response.
	StatusUnauthorized = 401

	// PaymentRequired - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.5.2
	// This response code is reserved for future use.
	StatusPaymentRequired = 402

	// Forbidden - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.5.3
	// The client does not have access rights to the content.
	StatusForbidden = 403

	// NotFound - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.5.4
	// The server can not find requested resource.
	StatusNotFound = 404

	// MethodNotAllowed - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.5.5
	// The request method is known by the server but has been disabled and cannot be used.
	StatusMethodNotAllowed = 405

	// NotAcceptable - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.5.6
	// The server doesn't find any content following the criteria given by the user agent.
	StatusNotAcceptable = 406

	// ProxyAuthenticationRequired - Official Documentation @ https://tools.ietf.org/html/rfc7235#section-3.2
	// Authentication is needed to be done by a proxy.
	StatusProxyAuthenticationRequired = 407

	// RequestTimeout - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.5.7
	// This response is sent on an idle connection by some servers, even without any previous request by the client.
	StatusRequestTimeout = 408

	// Conflict - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.5.8
	// This response is sent when a request conflicts with the current state of the server.
	StatusConflict = 409

	// Gone - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.5.9
	// This response would be sent when the requested content has been permanently deleted from server, with no forwarding address.
	StatusGone = 410

	// LengthRequired - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.5.10
	// The server rejected the request because the Content-Length header field is not defined and the server requires it.
	StatusLengthRequired = 411

	// PreconditionFailed - Official Documentation @ https://tools.ietf.org/html/rfc7232#section-4.2
	// The client has indicated preconditions in its headers which the server does not meet.
	StatusPreconditionFailed = 412

	// RequestTooLong - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.5.11
	// Request entity is larger than limits defined by server.
	StatusRequestTooLong = 413

	// RequestURITooLong - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.5.12
	// The URI requested by the client is longer than the server is willing to interpret.
	StatusRequestURITooLong = 414

	// UnsupportedMediaType - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.5.13
	// The media format of the requested data is not supported by the server.
	StatusUnsupportedMediaType = 415

	// RequestedRangeNotSatisfiable - Official Documentation @ https://tools.ietf.org/html/rfc7233#section-4.4
	// The range specified by the Range header field in the request can't be fulfilled.
	StatusRequestedRangeNotSatisfiable = 416

	// ExpectationFailed - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.5.14
	// The expectation indicated by the Expect request header field can't be met by the server.
	StatusExpectationFailed = 417

	// ImATeapot - Official Documentation @ https://tools.ietf.org/html/rfc2324#section-2.3.2
	// Any attempt to brew coffee with a teapot should result in the error code "418 I'm a teapot".
	StatusImATeapot = 418

	// InsufficientSpaceOnResource - Official Documentation @ https://tools.ietf.org/html/rfc2518#section-10.6
	// The method could not be performed on the resource because the server is unable to store the representation needed.
	StatusInsufficientSpaceOnResource = 419

	// MethodFailure - Deprecated
	// A deprecated response used by the Spring Framework when a method has failed.
	StatusMethodFailure = 420

	// MisdirectedRequest - Official Documentation @ https://datatracker.ietf.org/doc/html/rfc7540#section-9.1.2
	// A server is not able to produce a response for the combination of scheme and authority in the request URI.
	StatusMisdirectedRequest = 421

	// UnprocessableEntity - Official Documentation @ https://tools.ietf.org/html/rfc2518#section-10.3
	// The request was well-formed but was unable to be followed due to semantic errors.
	StatusUnprocessableEntity = 422

	// Locked - Official Documentation @ https://tools.ietf.org/html/rfc2518#section-10.4
	// The resource that is being accessed is locked.
	StatusLocked = 423

	// FailedDependency - Official Documentation @ https://tools.ietf.org/html/rfc2518#section-10.5
	// The request failed due to failure of a previous request.
	StatusFailedDependency = 424

	// PreconditionRequired - Official Documentation @ https://tools.ietf.org/html/rfc6585#section-3
	// The origin server requires the request to be conditional.
	StatusPreconditionRequired = 428

	// TooManyRequests - Official Documentation @ https://tools.ietf.org/html/rfc6585#section-4
	// The user has sent too many requests in a given amount of time ("rate limiting").
	StatusTooManyRequests = 429

	// RequestHeaderFieldsTooLarge - Official Documentation @ https://tools.ietf.org/html/rfc6585#section-5
	// The server is unwilling to process the request because its header fields are too large.
	StatusRequestHeaderFieldsTooLarge = 431

	// UnavailableForLegalReasons - Official Documentation @ https://tools.ietf.org/html/rfc7725
	// The user-agent requested a resource that cannot legally be provided.
	StatusUnavailableForLegalReasons = 451

	// InternalServerError - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.6.1
	// The server encountered an unexpected condition that prevented it from fulfilling the request.
	StatusInternalServerError = 500

	// NotImplemented - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.6.2
	// The request method is not supported by the server and cannot be handled.
	StatusNotImplemented = 501

	// BadGateway - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.6.3
	// The server, while working as a gateway, got an invalid response.
	StatusBadGateway = 502

	// ServiceUnavailable - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.6.4
	// The server is not ready to handle the request.
	StatusServiceUnavailable = 503

	// GatewayTimeout - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.6.5
	// The server is acting as a gateway and cannot get a response in time.
	StatusGatewayTimeout = 504

	// HTTPVersionNotSupported - Official Documentation @ https://tools.ietf.org/html/rfc7231#section-6.6.6
	// The HTTP version used in the request is not supported by the server.
	StatusHTTPVersionNotSupported = 505

	// InsufficientStorage - Official Documentation @ https://tools.ietf.org/html/rfc2518#section-10.6
	// The server is unable to store the representation needed to successfully complete the request.
	StatusInsufficientStorage = 507

	// NetworkAuthenticationRequired - Official Documentation @ https://tools.ietf.org/html/rfc6585#section-6
	// The client needs to authenticate to gain network access.
	StatusNetworkAuthenticationRequired = 511
)

var msg = map[int]string{
	StatusContinue:                      "Continue",
	StatusSwitchingProtocols:            "Switching Protocols",
	StatusProcessing:                    "Processing",
	StatusOK:                            "OK",
	StatusCreated:                       "Created",
	StatusAccepted:                      "Accepted",
	StatusNonAuthoritativeInformation:   "Non Authoritative Information",
	StatusNoContent:                     "No Content",
	StatusResetContent:                  "Reset Content",
	StatusPartialContent:                "Partial Content",
	StatusMultiStatus:                   "Multi-Status",
	StatusMultipleChoices:               "Multiple Choices",
	StatusMovedPermanently:              "Moved Permanently",
	StatusMovedTemporarily:              "Moved Temporarily",
	StatusSeeOther:                      "See Other",
	StatusNotModified:                   "Not Modified",
	StatusUseProxy:                      "Use Proxy",
	StatusTemporaryRedirect:             "Temporary Redirect",
	StatusPermanentRedirect:             "Permanent Redirect",
	StatusBadRequest:                    "Bad Request",
	StatusUnauthorized:                  "Unauthorized",
	StatusPaymentRequired:               "Payment Required",
	StatusForbidden:                     "Forbidden",
	StatusNotFound:                      "Not Found",
	StatusMethodNotAllowed:              "Method Not Allowed",
	StatusNotAcceptable:                 "Not Acceptable",
	StatusProxyAuthenticationRequired:   "Proxy Authentication Required",
	StatusRequestTimeout:                "Request Timeout",
	StatusConflict:                      "Conflict",
	StatusGone:                          "Gone",
	StatusLengthRequired:                "Length Required",
	StatusPreconditionFailed:            "Precondition Failed",
	StatusRequestTooLong:                "Request Entity Too Large",
	StatusRequestURITooLong:             "Request-URI Too Long",
	StatusUnsupportedMediaType:          "Unsupported Media Type",
	StatusRequestedRangeNotSatisfiable:  "Requested Range Not Satisfiable",
	StatusExpectationFailed:             "Expectation Failed",
	StatusImATeapot:                     "I'm a teapot",
	StatusInsufficientSpaceOnResource:   "Insufficient Space on Resource",
	StatusMethodFailure:                 "Method Failure",
	StatusMisdirectedRequest:            "Misdirected Request",
	StatusUnprocessableEntity:           "Unprocessable Entity",
	StatusLocked:                        "Locked",
	StatusFailedDependency:              "Failed Dependency",
	StatusPreconditionRequired:          "Precondition Required",
	StatusTooManyRequests:               "Too Many Requests",
	StatusRequestHeaderFieldsTooLarge:   "Request Header Fields Too Large",
	StatusUnavailableForLegalReasons:    "Unavailable For Legal Reasons",
	StatusInternalServerError:           "Internal Server Error",
	StatusNotImplemented:                "Not Implemented",
	StatusBadGateway:                    "Bad Gateway",
	StatusServiceUnavailable:            "Service Unavailable",
	StatusGatewayTimeout:                "Gateway Timeout",
	StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",
	StatusInsufficientStorage:           "Insufficient Storage",
	StatusNetworkAuthenticationRequired: "Network Authentication Required",
}
