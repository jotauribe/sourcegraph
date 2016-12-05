import * as classNames from "classnames";
import * as React from "react";
import * as base from "sourcegraph/components/styles/_base.css";
import * as styles from "sourcegraph/components/styles/select.css";
import { Alert } from "sourcegraph/components/symbols";
import { ChevronDown } from "sourcegraph/components/symbols/Zondicons";

interface Props {
	block?: boolean;
	className?: string;
	children?: any;
	label?: string;
	placeholder?: string;
	helperText?: string;
	error?: boolean;
	errorText?: string;
	style?: any;
	defaultValue?: string;
}

type State = any;

export class Select extends React.Component<Props, State> {
	static defaultProps: any = {
		block: true,
	};

	render(): JSX.Element | null {
		const {style, block, className, placeholder, label, helperText, error, errorText, children, defaultValue} = this.props;
		return (
			<div className={className}>
				{label && <div>{label} <br /></div>}
				<select
					required={true}
					style={style}
					defaultValue={defaultValue}
					className={classNames(styles.select, error ? styles.border_red : styles.border_neutral, block ? styles.block : null)}
					placeholder={placeholder ? placeholder : ""}>
					{children}
				</select>
				<ChevronDown style={{ marginLeft: "-28px" }} width={11} className={styles.icon} />
				{helperText && <em className={classNames(styles.small, styles.block, base.mt2)}>{helperText}</em>}
				{errorText &&
					<div className={classNames(styles.red, base.mv2)}>
						<Alert width={16} className={classNames(base.mr2, styles.red_fill)} style={{ marginTop: "-4px" }} />
						This is an error message.
					</div>
				}
			</div>
		);
	}
}
