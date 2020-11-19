import React, { FunctionComponent } from 'react'

export interface ExecutorLogsEntryProps {
    entry: { command: string[]; out: string }
    className?: string
}

export const ExecutorLogsEntry: FunctionComponent<ExecutorLogsEntryProps> = ({ entry, className }) => (
    <>
        <div className={className}>
            {/* TODO - better styles */}
            <strong>{entry.command.join(' ')}</strong>

            <pre className="bg-code rounded p-3">
                {/* TODO - don't do danger here */}
                <code dangerouslySetInnerHTML={{ __html: entry.out }} />
            </pre>
        </div>
    </>
)
