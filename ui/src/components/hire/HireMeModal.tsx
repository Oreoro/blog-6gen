/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import React from 'react';
import { Modal, Button } from 'react-bootstrap';
import { useTranslation } from 'react-i18next';

import type { PublicHireInfoResp } from '@/common/interface';

interface Props {
  show: boolean;
  onHide: () => void;
  data?: PublicHireInfoResp;
}

const HireMeModal: React.FC<Props> = ({ show, onHide, data }) => {
  const { t } = useTranslation('translation');
  const rate = data?.rate;
  const contact = data?.contact;
  const formatRate = () => {
    if (!rate) return '';
    const unitKey =
      rate.unit === 'day'
        ? 'per_day'
        : rate.unit === 'project'
          ? 'per_project'
          : 'per_hour';
    return `${rate.currency || ''}${rate.amount ?? ''} ${t(`hire.${unitKey}`)}
    `;
  };
  return (
    <Modal show={show} onHide={onHide} centered>
      <Modal.Header closeButton>
        <Modal.Title>{t('hire.hire_me')}</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        {rate && (
          <div className="mb-3">
            <strong>{t('hire.rate')}:</strong> {formatRate()}
          </div>
        )}
        {contact?.email && (
          <div className="mb-2">
            <strong>{t('hire.contact_via_email')}:</strong>{' '}
            <a href={`mailto:${contact.email}`}>{contact.email}</a>
          </div>
        )}
        {contact?.url && (
          <div className="mb-2">
            <strong>{t('hire.open_website')}:</strong>{' '}
            <a href={contact.url} target="_blank" rel="noreferrer">
              {contact.url}
            </a>
          </div>
        )}
        {data?.note && (
          <div>
            <strong>{t('hire.note')}:</strong>
            <div className="mt-1">{data.note}</div>
          </div>
        )}
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={onHide}>
          Close
        </Button>
      </Modal.Footer>
    </Modal>
  );
};

export default HireMeModal;
