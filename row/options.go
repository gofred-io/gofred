package row

import "github.com/gofred-io/gofred/style"

type Options func(row *Row)

func ID(id string) Options {
	return func(row *Row) {
		row.SetID(id)
	}
}

func CrossAxisAlignment(crossAxisAlignment style.AlignItemsType) Options {
	return func(row *Row) {
		if row.style.Display == nil {
			row.style.Display = &style.Display{}
		}

		row.style.Display.AlignItems = crossAxisAlignment
	}
}

func MainAxisAlignment(mainAxisAlignment style.JustifyContentType) Options {
	return func(row *Row) {
		if row.style.Display == nil {
			row.style.Display = &style.Display{}
		}

		row.style.Display.JustifyContent = mainAxisAlignment
	}
}

func Gap(columnGap int) Options {
	return func(row *Row) {
		row.style.Display.ColumnGap = columnGap
	}
}
