package serrors

import (
	"context"

	"github.com/nicksnyder/go-i18n/v2/i18n"

	"go-template/internal/pkg/ccontext"
	"go-template/pkg/container"
	"go-template/pkg/errors"
)

func WrapI18n(ctx context.Context, code errors.Code, message string, errs ...error) errors.IError {
	lang := ccontext.GetLang(ctx)
	var bundle, _ = container.Resolver[*i18n.Bundle]()
	if bundle == nil {
		return errors.ErrorTraceCtx(ctx, code, message, errs...)
	}

	localizer := i18n.NewLocalizer(bundle, lang)

	msg, err := localizer.Localize(
		&i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID:    message,
				Other: message,
			},
		},
	)
	if err != nil {
		msg = message
	}
	return errors.ErrorTraceCtx(ctx, code, msg, errs...)
}

func ErrUnauthenticated(ctx context.Context, err error) errors.IError {
	return WrapI18n(ctx, errors.Unauthenticated, "Unauthenticated", err)
}

func ErrInternal(ctx context.Context, err error) errors.IError {
	return WrapI18n(ctx, errors.Internal, "Internal error", err)
}

func ErrFailedPrecondition(ctx context.Context, err error) errors.IError {
	return WrapI18n(ctx, errors.FailedPrecondition, "FailedPrecondition", err)
}

func ErrInvalidArgument(ctx context.Context, err error) errors.IError {
	return WrapI18n(ctx, errors.InvalidArgument, "InvalidArgument", err)
}
