import React, { FunctionComponent } from 'react'
import { LsifIndexFields } from '../../../graphql-operations'
import { ExecutorLogsEntry } from './ExecutorLogsEntry'

export interface ExecutorLogsProps {
    index: LsifIndexFields
    className?: string
}

export const ExecutorLogs: FunctionComponent<ExecutorLogsProps> = ({ index, className }) => (
    <>
        <h3>Output logs</h3>
        {index.logContents.length === 0 ? (
            <>No output logs</>
        ) : (
            index.logContents.map(entry => (
                <ExecutorLogsEntry key={`${entry.command.join(' ')}${entry.out}`} entry={entry} className={className} />
            ))
        )}
    </>
)
