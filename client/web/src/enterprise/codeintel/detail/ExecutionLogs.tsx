import React, { FunctionComponent } from 'react'
import { LsifIndexFields } from '../../../graphql-operations'
import { ExecutionLogEntry } from './ExecutionLogEntry'

export interface ExecutionLogsProps {
    index: LsifIndexFields
    className?: string
}

export const ExecutionLogs: FunctionComponent<ExecutionLogsProps> = ({ index, className }) => (
    <>
        <h3>Output logs</h3>
        {index.executionLogs.length === 0 ? (
            <>No output logs</>
        ) : (
            index.executionLogs.map(entry => (
                <ExecutionLogEntry key={`${entry.command.join(' ')}${entry.out}`} entry={entry} className={className} />
            ))
        )}
    </>
)
